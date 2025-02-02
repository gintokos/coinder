import CoinCard from '../../components/coincard/СoinСard'
import classes from './browsing.module.css'
import { useEffect, useState, useCallback } from 'react'
import ScrollBtn from '../../components/scrollbtn/ScrollBtn'

export default function Browsing() {
   const [currentIndex, setCurrentIndex] = useState(0);
   const [touchStart, setTouchStart] = useState(null);
   const totalCards = 5;

   const handleScrollUp = useCallback(() => {
       setCurrentIndex(prev => Math.max(prev - 1, 0));
   }, [])

   const handleScrollDown = useCallback(() => {
       setCurrentIndex(prev => Math.min(prev + 1, totalCards - 1));
   }, [totalCards])

   useEffect(() => {
       const handleWheel = (e) => {
           e.preventDefault();

           if (e.deltaY > 0) {
               handleScrollDown();
           } else {
               handleScrollUp();
           }
       };

       const handleTouchStart = (e) => {
           setTouchStart(e.touches[0].clientY);
       };

       const handleTouchMove = (e) => {
           e.preventDefault();
           if (!touchStart) return;

           const touchEnd = e.touches[0].clientY;
           const diff = touchStart - touchEnd;

           if (Math.abs(diff) > 50) {
               if (diff > 0) {
                   handleScrollDown();
               } else {
                   handleScrollUp();
               }
               setTouchStart(null);
           }
       };

       const handleTouchEnd = () => {
           setTouchStart(null);
       };

       const container = document.querySelector(`.${classes.container}`);
       
       container.addEventListener('wheel', handleWheel, { passive: false });
       container.addEventListener('touchstart', handleTouchStart, { passive: false });
       container.addEventListener('touchmove', handleTouchMove, { passive: false });
       container.addEventListener('touchend', handleTouchEnd);

       return () => {
           container.removeEventListener('wheel', handleWheel);
           container.removeEventListener('touchstart', handleTouchStart);
           container.removeEventListener('touchmove', handleTouchMove);
           container.removeEventListener('touchend', handleTouchEnd);
       };
   }, [touchStart, handleScrollDown, handleScrollUp]);

   useEffect(() => {
       const cards = document.querySelectorAll(`.${classes.card}`);
       cards[currentIndex]?.scrollIntoView({
           behavior: 'smooth',
           block: 'start'
       });
   }, [currentIndex]);

   return (
       <>
           <div className={classes.container}>
               <ScrollBtn 
                   onClick={handleScrollUp} 
                   className={`${classes.btn} ${classes.up}`} 
               />
               <ScrollBtn 
                   onClick={handleScrollDown} 
                   className={`${classes.btn} ${classes.down}`} 
               />
               <CoinCard coin={1} className={classes.card} />
               <CoinCard coin={2} className={classes.card} />
               <CoinCard coin={3} className={classes.card} />
               <CoinCard coin={4} className={classes.card} />
               <CoinCard coin={5} className={classes.card} />
           </div>
       </>
   );
}