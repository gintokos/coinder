import { createAsyncThunk } from "@reduxjs/toolkit"
import { coinderApi } from '../api/api'

const BOT_ID = import.meta.env.VITE_BOT_ID

export const checkAuth = createAsyncThunk(
    'auth/checkAuth',
    async (_, { rejectWithValue }) => {
        try {
            if (window.Telegram.WebApp.initData !== ""){
                const tg = window.Telegram.WebApp

                await new Promise(resolve => {
                    tg.ready()
                    resolve()
                })

                const params = new URLSearchParams(tg.initData);
                const userData = JSON.parse(params.get('user'));
                
                const data = {
                    id: parseInt(userData.id), 
                    first_name: userData.first_name,
                    username: userData.username,
                    photo_url: userData.photo_url,
                    auth_date: parseInt(params.get('auth_date')),
                    hash: params.get('hash')
                };

                const entries = Array.from(params.entries());
                const sortedEntries = entries
                    .filter(([key]) => key !== 'hash')
                    .sort(([a], [b]) => a.localeCompare(b))
                    .map(([key, value]) => `${key}=${value}`);

                const checkString = sortedEntries.join('\n');

                const response = await coinderApi.login({
                    datastr: checkString,
                    data: data,
                })
                return {
                    isAuth: true,
                    id: data.id,
                    firstName: data.first_name,
                    username: data.username,
                    photoUrl: data.photo_url,
                    authDate: data.auth_date,
                };
            }

            const ok = await coinderApi.ping()
            if (ok) {
                const user = {
                    id: localStorage.getItem('id'),
                    username: localStorage.getItem('username'),
                    firstName: localStorage.getItem('firstName'),
                    photoUrl: localStorage.getItem('photoUrl'),
                    authDate: localStorage.getItem('authDate')
                }
                coinderApi.refresh()
                return {
                    isAuth: true,
                    ...user
                }
            }
        } catch (error) {
            return rejectWithValue(error.message)
        }
    }
)

export const webAuth = createAsyncThunk(
    'auth/webAuth',
    async(_, { rejectWithValue }) => {
        try {
            let user = {}
            await new Promise((resolve, reject) => {
                window.Telegram.Login.auth(
                    {
                        bot_id: BOT_ID,
                        request_access: true,
                        embed: false
                    },
                    async (data) => {
                        if (data) {
                            const { hash, ...rest } = data
                            const dataArr = Object.entries(rest).map(([key, value]) => `${key}=${value}`).sort()
                            const checkString = dataArr.join('\n') 
                            
                            user = {
                                isAuth: true,
                                id: rest.id,
                                username: rest.username,
                                firstName: rest.first_name,
                                photoUrl: rest.photo_url,
                                authDate: rest.auth_date
                            }

                            Object.entries(user).forEach(([key, value]) => {
                                localStorage.setItem(key, value)
                            })

                            await coinderApi.login({
                                datastr: checkString,
                                data: data,
                            })
                        } else {
                            rejectWithValue(new Error("auth failed"))
                        }
                        resolve()
                    }
                )
            })

            return { ...user }

        } catch (error) {
            return rejectWithValue(error)
        }
    }
)