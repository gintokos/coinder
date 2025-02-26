import React, { useState, useEffect } from 'react';
import { Outlet } from 'react-router';
import Footer from '../components/footer/footer';
import classes from './Page.module.css';
import useTelegramBackButton from '../hooks/back';
import Card from '../components/card/card';

const Page = () => {
  useTelegramBackButton();

  const [isLandscape, setIsLandscape] = useState(false);
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const checkOrientationAndDevice = () => {
      const landscapeMode = window.matchMedia("(orientation: landscape)").matches;
      const mobileDevice = window.matchMedia("(max-width: 768px)").matches;

      console.log('Checking: landscape=', landscapeMode, 'mobile=', mobileDevice); // Отладка

      setIsLandscape(landscapeMode);
      setIsMobile(mobileDevice);
    };

    checkOrientationAndDevice();

    const orientationQuery = window.matchMedia("(orientation: landscape)");
    const mobileQuery = window.matchMedia("(max-width: 768px)");

    const orientationHandler = (e) => {
      console.log('Orientation changed: landscape=', e.matches); // Отладка
      setIsLandscape(e.matches);
    };

    const deviceHandler = (e) => {
      console.log('Device changed: mobile=', e.matches); // Отладка
      setIsMobile(e.matches);
    };

    orientationQuery.addEventListener('change', orientationHandler);
    mobileQuery.addEventListener('change', deviceHandler);

    return () => {
      orientationQuery.removeEventListener('change', orientationHandler);
      mobileQuery.removeEventListener('change', deviceHandler);
    };
  }, []);

  const showOrientationWarning = isMobile && isLandscape;
  console.log('Rendering: showWarning=', showOrientationWarning, 'isMobile=', isMobile, 'isLandscape=', isLandscape); // Отладка

  return (
    <div className={classes.layout}>
      <main className={classes.content}>
        {showOrientationWarning ? (
          <Card>
            <h2 style={{ textAlign: 'center', fontSize: '2rem', color: 'var(--accent-blue)' }}>
              Please turn your device to portrait mode
            </h2>
          </Card>
        ) : (
          <Outlet />
        )}
      </main>
      <footer className={classes.footer}>
        <Footer />
      </footer>
    </div>
  );
};

export default Page;