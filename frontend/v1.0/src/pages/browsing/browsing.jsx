import CoinCard from '../../components/coincard/СoinСard';
import classes from './browsing.module.css';
import { useState, useEffect } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Navigation, Mousewheel, Keyboard, Manipulation } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';
import ScrollBtn from '../../components/scrollbtn/scrollbtn';

export default function Browsing() {
    const [swiperRef, setSwiperRef] = useState(null);
    const [slides, setSlides] = useState([1, 2, 3, 4, 5]);
    const [currentIndex, setCurrentIndex] = useState(0);

    const onReachEnd = () => {
        if (swiperRef) {
            const index = swiperRef.activeIndex;
            setCurrentIndex(index);
            console.log("end");
            
            setSlides(prev => [...prev, prev.length + 1]);
        }
    };

        useEffect(() => {
        if (swiperRef && currentIndex !== 0) {
            swiperRef.slideTo(currentIndex + 1, 0, false);
        }
    }, [slides.length]);

    const isMobile = window.innerWidth <= 680;

    return (
        <div className={classes.container}>
            <Swiper
                onSwiper={setSwiperRef}
                modules={[Navigation, Mousewheel, Keyboard, Manipulation]}
                direction="vertical"
                slidesPerView={1}
                spaceBetween={30}
                centeredSlides={true}
                mousewheel={true}
                keyboard={{
                    enabled: true,
                }}
                breakpoints={{
                    680: {
                        navigation: {
                            enabled: true,
                        },
                    }
                }}
                onReachEnd={onReachEnd}
                onSlideChange={(swiper) => {
                    console.log('Текущий слайд:', swiper.activeIndex);
                }}
            >
                {slides.map((coinValue) => (
                    <SwiperSlide key={coinValue}>
                        <CoinCard coin={coinValue} />
                    </SwiperSlide>
                ))}
            </Swiper>
            
            {!isMobile && (
                <>
                    <ScrollBtn 
                        onClick={() => swiperRef?.slidePrev()} 
                        className={`${classes.btn} ${classes.up}`} 
                    />
                    <ScrollBtn 
                        onClick={() => swiperRef?.slideNext()} 
                        className={`${classes.btn} ${classes.down}`} 
                    />
                </>
            )}
        </div>
    );
}