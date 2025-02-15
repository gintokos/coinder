import { useLocation, Navigate } from "react-router"
import store from "../redux/store"

import { Outlet } from "react-router"


export function RequireAuth() {
    const location = useLocation()
    const auth = store.getState().auth

    if (auth.isAuth) {
        return (
            <Outlet/>
        )
    } else {
        return <Navigate to="/auth" state={{from:location}} />
    }
}