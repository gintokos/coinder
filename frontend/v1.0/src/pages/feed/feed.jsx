import CoinCard from '../../components/coincard/СoinСard';
import classes from './feed.module.css';
import { useState, useEffect } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Navigation, Mousewheel, Keyboard, Manipulation } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';
import ScrollBtn from '../../components/scrollbtn/scrollbtn';
import ReactDOM from 'react-dom/client';
import { useLoaderData } from 'react-router';
import { coinderApi } from '../../api/api';

export default function Feed() {
    console.log("Feed render")
    let [coins, params] = useLoaderData()

    const [swiperRef, setSwiperRef] = useState(null);

    const appendNewSlides = async () => {
        if (swiperRef) {
            try {
                const response = await coinderApi.coins();
                if (response && response.data && Array.isArray(response.data)) {
                    response.data.forEach(coin => {
                        const slideContainer = document.createElement('div');
                        slideContainer.className = 'swiper-slide';
                        
                        swiperRef.appendSlide(slideContainer);
                        
                        const root = ReactDOM.createRoot(slideContainer);
                        root.render(<CoinCard coin={coin} />);
                    });
                }
            } catch (error) {
                console.error('Ошибка при загрузке новых монет:', error);
            }
        }
    }

    useEffect(() => {
        let isThrottled = false;
        const THROTTLE_DELAY = 300;
    
        const handleKeyboard = (event) => {
            if (isThrottled) return;
    
            if (event.key === 'ArrowUp') {
                console.log('up');
                swiperRef?.slidePrev();
                isThrottled = true;
            }
            if (event.key === 'ArrowDown') {
                console.log('down');
                swiperRef?.slideNext();
                isThrottled = true;
            }
    
            setTimeout(() => {
                isThrottled = false;
            }, THROTTLE_DELAY);
        }
    
        document.addEventListener('keydown', handleKeyboard);
    
        return () => {
            document.removeEventListener('keydown', handleKeyboard);
        };
    }, [swiperRef]);

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
                breakpoints={{
                    680: {
                        navigation: {
                            enabled: true,
                        },
                    }
                }}
                onKeyDownCapture={()=>console.log("up")}
                onSlideChange={(swiper) => {
                    console.log('Текущий слайд:', swiper.activeIndex);
                    if (swiper.activeIndex === swiper.slides.length - 10) {
                        console.log("appending")
                        appendNewSlides();
                        swiper.removeSlide(Array.from({length: 70}, (_, i) => i))
                    }
                }}
            >
                {coins.map((coin) => (
                    <SwiperSlide key={coin.id}>
                        <CoinCard coin={coin} />
                    </SwiperSlide>
                ))}
            </Swiper>

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
        </div>
    );
}

const feedLoader = async ({request, params}) => {
    const response = await coinderApi.coins()
    console.log('response: ', response)
    return [response.data, params]
}

export { feedLoader }