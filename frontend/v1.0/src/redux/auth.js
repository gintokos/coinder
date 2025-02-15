import { createSlice } from "@reduxjs/toolkit"
import { checkAuth, webAuth } from "../thunk/auth.js"

const initialState = {
    isAuth: false,
    id: 0,
    username: '',
    firstName: '',
    photoUrl: '',
    authDate: '',
    loading: false,
    checkError: null,
    webAuthError: null,
}

export const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        login: (state, action) => {
            Object.assign(state, action.payload)
        },
        logout: (state) => {
            console.log("dispatch logout")
            Object.assign(state, initialState)
        }
    },
    extraReducers: (builder) => {
        builder
            .addCase(checkAuth.pending, (state) => {
                state.loading = true;
                state.checkError = null;
            })
            .addCase(checkAuth.fulfilled, (state, action) => {
                Object.assign(state, action.payload);
                state.loading = false;
            })
            .addCase(checkAuth.rejected, (state, action) => {
                state.loading = false;
                state.checkError = action.payload;
                state.isAuth = false;
            })
            .addCase(webAuth.pending, (state) => {
                state.loading = true;
                state.webAuthError = null;
            })
            .addCase(webAuth.fulfilled, (state, action) => {
                Object.assign(state, action.payload);
                state.loading = false;
            })
            .addCase(webAuth.rejected, (state, action) => {
                state.loading = false;
                state.webAuthError = action.payload;
                state.isAuth = false;
            });
    }
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;