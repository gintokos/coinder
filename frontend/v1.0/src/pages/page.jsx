import { Layout } from 'antd'
import Header from '../components/header/header'
import Footer from '../components/footer/footer'
import classes from './page.module.css'
import { Outlet } from 'react-router'


export default function Page() {
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