const HTTP_METHODS = {
    GET: 'GET',
    POST: 'POST',
    PUT: 'PUT',
    DELETE: 'DELETE',
    PATCH: 'PATCH'
};

const API_SOURCES = {
    WEB: "web",
    MINIAPP: "miniapp"
};

const DEFAULT_TIMEOUT = 5000; 

class ApiError extends Error {
    constructor(message, status, data = null) {
        super(message);
        this.name = 'ApiError';
        this.status = status;
        this.data = data;
    }
}

class CoinderApi {
    #baseurl;
    #timeout;

    constructor(baseurl, timeout = DEFAULT_TIMEOUT) {
        console.log(baseurl)
        this.#baseurl = baseurl;
        this.#timeout = timeout;
        this.dispatch = undefined
    }

    #getSource = () => {
        if (window?.Telegram?.WebApp?.initData !== undefined && 
            window.Telegram.WebApp.initData !== "") {
            return API_SOURCES.MINIAPP;
        }
        return API_SOURCES.WEB;
    }

    #handleResponse = async (response) => {
        if (!response.ok) {
            const errorData = await response.json().catch(() => null);
            switch (response.status) {
                case 401:
                    this.dispatch({ type: 'auth/logout' })
                    throw new ApiError('Unauthorized', 401, errorData);
                case 403:
                    throw new ApiError('Forbidden', 403, errorData);
                case 404:
                    throw new ApiError('Not Found', 404, errorData);
                case 429:
                    throw new ApiError('Too Many Requests', 429, errorData);
                default:
                    throw new ApiError(
                        `HTTP error! status: ${response.status}`,
                        response.status,
                        errorData
                    );
            }
        }

        return response.json();
    }

    fetch = async (endpoint, options = {}) => {
        const controller = new AbortController();
        const timeout = setTimeout(() => controller.abort(), this.#timeout);

        try {
            const body = options.body ? JSON.stringify(options.body) : undefined;

            const defaultOptions = {
                method: HTTP_METHODS.GET,
                headers: {
                    'Content-Type': 'application/json',
                    'T-Source-H': this.#getSource(),
                },
                credentials: 'include',
                signal: controller.signal
            };

            const fetchOptions = {
                ...defaultOptions,
                ...options,
                body,
                headers: {
                    ...defaultOptions.headers,
                    ...(options.headers || {}),
                },
            };

            const response = await fetch(`${this.#baseurl}${endpoint}`, fetchOptions);
            clearTimeout(timeout);
            
            return await this.#handleResponse(response);

        } catch (error) {
            clearTimeout(timeout);

            if (error instanceof ApiError) {
                throw error;
            }

            if (error.name === 'AbortError') {
                throw new ApiError('Request timeout', 408);
            }

            if (error instanceof TypeError) {
                console.error("Network error:", error);
                throw new ApiError('Network error', 0);
            }

            console.error("API request error:", error);
            throw error;
        }
    }

    ping = async () => {
        try {
            const result = await this.fetch("/ping");
            return result.pong === "success";
        } catch (error) {
            console.error("Error on ping:", error);
            return false;
        }
    }

    refresh = async () => {
        try {
            return await this.fetch("/auth/refresh", { 
                method: HTTP_METHODS.GET 
            });
        } catch (error) {
            console.error("Error on refresh:", error);
            throw error;
        }
    }

    login = async (credentials) => {
        const dataCheck = Object.entries(credentials.data)
            .filter(([key]) => key !== 'hash') 
            .sort(([a], [b]) => a.localeCompare(b))
            .map(([key, value]) => `${key}=${value}`)
            .join('\n');
    
        await this.fetch("/auth", {
            method: HTTP_METHODS.POST,
            body: {...credentials}
        });
        
        return await this.updateUser()
    }

    logout = async () => {
        return await this.fetch("/auth/logout", {
            method: HTTP_METHODS.POST
        });
    }
    
    updateUser = async () => {
        try {
            await this.fetch("/user/update", {
                method: HTTP_METHODS.POST
            })
        } catch (error) {
            console.error("error no updating user")
        }
    }

    coins = async (sOptions) => {
        const defaultOpt = {
            limit: 100,
            page: 1,
            sorted_by: "BY_PRICE",
            liked_by_user: true,
            user_id_target: null,
            liked_today: false
        }

        sOptions = {...defaultOpt, ...sOptions}
    
        try {
            return await this.fetch("/coins/default", {
                method: HTTP_METHODS.POST,
                body: sOptions
            })
        } catch (error) {
            console.error("error on getting coins", error)
            throw error
        }
    }

    like = async (coinId) => {
        try {
            return await this.fetch("/coins/like", {
                method: HTTP_METHODS.POST,
                body: { coin_id: coinId }
            })
        } catch (error) {
            console.error("error on liking coin", error)
        }
    }

    dislike = async (coinId) => {
        try {
            return await this.fetch("/coins/dislike", {
                method: HTTP_METHODS.POST,
                body: { coin_id: coinId }
            })
        } catch (error) {
            console.error("error on disliking coin", error)
        }
    }

    createInvoice = async(amount) => {
        try {
            return await this.fetch("/user/create_invoice?amount=" + amount, {
                method: HTTP_METHODS.GET,
            })
        } catch (error) {
            console.error("error on creating invoice", error)
            throw error
        }
     }
}

const BASE_URL = import.meta.env.VITE_DOMAIN
export const coinderApi = new CoinderApi(BASE_URL)