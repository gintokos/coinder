import { Layout } from 'antd'
import Header from '../components/header/header'
import Footer from '../components/footer/footer'
import classes from './page.module.css'
import { Outlet } from 'react-router'
import { useSelector} from "react-redux"
import Auth from './auth/auth'

export default function Page() {
  const auth = useSelector((state) => state.auth)

  return (
    <Layout className={classes.layout}>
      <Layout.Header className={classes.header}>
        <Header />
      </Layout.Header>
      <Layout.Content className={classes.content}>
        <Outlet />
      </Layout.Content>
      <Layout.Footer className={classes.footer}>
        <Footer />
      </Layout.Footer>
    </Layout>
  );
}