import { useState } from 'react'

const useDoubleTap = (callback, delay = 300) => {
    const [lastTap, setLastTap] = useState(0);
  
    return {
      onTouchStart: (e) => {
        const now = new Date().getTime();
        if (now - lastTap < delay) {
          callback();
        }
        setLastTap(now);
      },
      onDoubleClick: callback
    };
};

export {useDoubleTap}