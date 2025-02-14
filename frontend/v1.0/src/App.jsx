import { Route, Routes } from "react-router";
import Page from "./pages/page.jsx"
import Browsing from "./pages/browsing/Browsing.jsx"
import { coinderApi } from "./api/api.js";


function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Page />}>
          <Route index element={<>YA GLAVNAYA</>} />
          <Route path="browsing" element={<Browsing />} />
        </Route>
      </Routes>
    </>
  );
}

export default App;