import { configureStore } from "@reduxjs/toolkit"
import authReducer from "./auth.js"
import { coinderApi } from "../api/api.js"

const store = configureStore({
    reducer: {
        auth: authReducer
    }
})

coinderApi.dispatch = store.dispatch

export default store