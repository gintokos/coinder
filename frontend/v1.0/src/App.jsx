import { Route, RouterProvider, createBrowserRouter, createRoutesFromElements } from "react-router";
import Page from "./pages/page.jsx"
import { useDispatch } from "react-redux";
import { useEffect } from "react";
import { checkAuth } from "./thunk/auth.js"
import Auth from "./pages/auth/auth.jsx";
import { RequireAuth } from "./hoc/requireauth.jsx"
import { NotFound } from "./pages/errorpage/notFound.jsx"
import Main from "./pages/main/main.jsx";
import Support from "./pages/support/support.jsx";
import Browsing from "./pages/browsing/browsing.jsx";
import Feed, { feedLoader } from "./pages/feed/feed.jsx";

const router = createBrowserRouter(createRoutesFromElements(
  <Route path="/" element={<Page />}>
    <Route index element={<Main/>} />
    <Route path="auth" element={<Auth/>}/>

    <Route element={<RequireAuth />}>
      <Route path="browsing" element={<Browsing />} />
      <Route path="browsing/feed" element={<Feed/>} loader={feedLoader} />
      <Route path="support_project" element={<Support/>} />
    </Route>

    <Route path="*" element={<NotFound/>} />
  </Route>  

))

function App() {
  const dispatch = useDispatch()
  useEffect(() => {
    dispatch(checkAuth())
  }, [dispatch])

  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;