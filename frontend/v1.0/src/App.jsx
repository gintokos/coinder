import { Route, Routes } from "react-router";
import Page from "./pages/page.jsx"
import Browsing from "./pages/browsing/Browsing.jsx"
import { useDispatch } from "react-redux";
import { useEffect } from "react";
import { checkAuth } from "./thunk/auth.js"
import { coinderApi } from "./api/api.js";
import { login } from "./redux/auth.js";

function App() {
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(checkAuth())
  }, [dispatch])

  return (
    <>
      <Routes>
        <Route />
        <Route path="/" element={<Page />}>
          <Route index element={<>YA GLAVNAYA</>} />
          <Route path="browsing" element={<Browsing />} />
        </Route>  
      </Routes>
    </>
  );
}

export default App;