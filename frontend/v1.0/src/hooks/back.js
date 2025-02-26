import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router';

const useTelegramBackButton = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const webApp = window.Telegram?.WebApp;

  useEffect(() => {
    if (!webApp) return;

    const handleBackClick = () => {
      if (location.pathname === '/') {
        webApp.close();
      } else {
        navigate(-1);
      }
    };

    if (location.pathname !== '/') {
      webApp.BackButton.show().onClick(handleBackClick);
    } else {
      webApp.BackButton.hide();
    }

    return () => {
      webApp.BackButton.offClick(handleBackClick);
    };
  }, [location.pathname, webApp]);
};

export default useTelegramBackButton;