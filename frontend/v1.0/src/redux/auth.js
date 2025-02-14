import { createSlice } from "@reduxjs/toolkit"

export const authSlice = createSlice({
    name: "auth", 
    initialState: {
        isAuth: false,
        id: 0,
        username: '',
        firstName: '',
        photoUrl: '',
        authDate: '',
    },
    reducers: {
        login: (state, action) => {
            Object.assign(state, action.payload)
        }
    }
})

export default authSlice.reducer