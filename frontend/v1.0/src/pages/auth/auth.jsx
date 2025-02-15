import { useDispatch, useSelector } from "react-redux"
import { webAuth } from "../../thunk/auth"


export default function Auth() {
    const dispatch = useDispatch()
    const { loading, webAuthError } = useSelector((state) => state.auth)

    const onClick = () => {
        dispatch(webAuth())
    }

    return (
        <>
            <button 
                onClick={onClick}
                disabled={loading}
            >
                {loading ? 'Загрузка...' : 'Авторизоваться'}
            </button>
            {webAuthError && (
                <div style={{color: 'red'}}>
                    Ошибка при авторизации: {webAuthError}
                </div>
            )}
        </>
    )
}