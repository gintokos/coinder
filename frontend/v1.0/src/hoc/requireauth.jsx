import { useLocation, Navigate } from "react-router"
import { useSelector } from 'react-redux'
import { Outlet } from "react-router"

export function RequireAuth() {
    const location = useLocation()
    const auth = useSelector(state => state.auth)

    if (auth.loading) {
        return <div>Loading...</div>
    }

    if (auth.isAuth) {
        return <Outlet/>
    } else {
        return <Navigate to="/auth" state={{from:location}} />
    }
}