import React from 'react';
import { Outlet } from 'react-router';
import Header from '../components/header/header';
import Footer from '../components/footer/footer';
import classes from './Page.module.css';

const Page = () => {
  return (
    <div className={classes.layout}>
      <main className={classes.content}>
        <Outlet />
      </main>
      
      <footer className={classes.footer}>
        <Footer />
      </footer>
    </div>
  );
};

export default Page;