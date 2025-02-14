import CoinCard from '../../components/coincard/СoinСard';
import classes from './browsing.module.css';
import { useState } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Navigation, Mousewheel, Keyboard, Manipulation } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';
import ScrollBtn from '../../components/scrollbtn/scrollbtn';
import ReactDOM from 'react-dom/client';

export default function Browsing() {
    const [swiperRef, setSwiperRef] = useState(null);
    let index = 5;

    const appendNewSlide = () => {
        if (swiperRef) {
            index += 1;
            
            const slideContainer = document.createElement('div');
            slideContainer.className = 'swiper-slide';
            
            swiperRef.appendSlide(slideContainer);
            
            const root = ReactDOM.createRoot(slideContainer);
            root.render(<CoinCard coin={index} />);
        }
    };

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
                onReachEnd={appendNewSlide}
                onSlideChange={(swiper) => {
                    console.log('Текущий слайд:', swiper.activeIndex);
                }}
            >
                {Array.from({length: 5}, (_, i) => i + 1).map((coinValue) => (
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