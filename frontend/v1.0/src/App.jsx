import { Route, Routes } from "react-router";
import Page from "./pages/page.jsx"
import Browsing from "./pages/browsing/Browsing.jsx"
import { useDispatch } from "react-redux";
import { useEffect } from "react";
import { checkAuth } from "./thunk/auth.js"
import Auth from "./pages/auth/auth.jsx";
import { RequireAuth } from "./hoc/requireauth.jsx"
import { NotFound } from "./pages/errorpage/notFound.jsx"

function App() {
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(checkAuth())
  }, [dispatch])

  return (
    <>
      <Routes>
        <Route path="/" element={<Page />}>
          <Route index element={<>YA GLAVNAYA</>} />
          <Route path="auth" element={<Auth/>}/>

          <Route path="p/" element={<RequireAuth/>} >
            <Route path="browsing" element={<Browsing />} />
          </Route>

          <Route path="*" element={<NotFound/>} />
        </Route>  
      </Routes>
    </>
  );
}

export default App;