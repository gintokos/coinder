import { useRef } from 'react';

const useDoubleTap = (callback, delay = 300) => {
  const lastTap = useRef(0); 

  return {
    onTouchStart: (e) => {
      const now = new Date().getTime();
      if (now - lastTap.current < delay) {
        callback();
      }
      lastTap.current = now;
    },
    onDoubleClick: callback,
  };
};

export { useDoubleTap };