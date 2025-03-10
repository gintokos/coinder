import CoinCard from '../../components/coincard/СoinСard';
import classes from './feed.module.css';
import { useState, useEffect, useCallback } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Navigation, Mousewheel, Keyboard, Manipulation } from 'swiper/modules';
import { useLocation } from 'react-router';
import 'swiper/css';
import 'swiper/css/navigation';
import ScrollBtn from '../../components/scrollbtn/Scrollbtn.jsx';
import ReactDOM from 'react-dom/client';
import { coinderApi } from '../../api/api';
import useCoins from '../../hooks/coins';
import Card from '../../components/card/card';

const LIMIT = 20

export default function Feed() {
    console.log("Feed render")
    const [swiperRef, setSwiperRef] = useState(null);
    const location = useLocation();
    
    const queryParams = new URLSearchParams(location.search);
    const sortedBy = queryParams.get('sorted_by') || 'BY_PRICE';
    const likedByUser = queryParams.get('liked_by_user') === 'true';
    
    const [nextCoins, firstCoins] = useCoins({
        limit: LIMIT
    });
    
    console.log("firstcoins: ", firstCoins);
    
    const appendNewSlides = useCallback(async () => {
        if (swiperRef) {
            try {
                const coins = await nextCoins();
                console.log("nextcoins: ", coins);
                
                if (coins && Array.isArray(coins)) {
                    coins.forEach(coin => {
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
    }, [swiperRef, nextCoins]);
    
    useEffect(() => {
        let isThrottled = false;
        const THROTTLE_DELAY = 100;
    
        const handleKeyboard = (event) => {
            if (isThrottled) return;
    
            if (event.key === 'ArrowUp') {
                swiperRef?.slidePrev();
                isThrottled = true;
            }
            if (event.key === 'ArrowDown') {
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

    if (firstCoins.length === 0) return (
        <Card>
            <h5 style={{fontSize: "2rem", color: "var(--accent-blue)", textAlign: "center", marginTop: "1.2rem"}}> Due this params no coins were found</h5>
        </Card>
    );

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
                onKeyDownCapture={() => console.log("up")}
                onSlideChange={(swiper) => {
                    console.log('Текущий слайд:', swiper.activeIndex);
                    if (swiper.activeIndex === 100) {
                        swiperRef.removeSlide(Array.from({length: 50}, (_, i) => i))
                    }
                    if (swiper.activeIndex === swiper.slides.length - 5) {
                        appendNewSlides()
                    }
                }}
            >
                {firstCoins.map((coin) => (
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

export const feedLoader = async ({request, params}) => {
    const url = new URL(request.url);
    
    const sortedBy = url.searchParams.get('sorted_by') || 'BY_PRICE';
    
    const likedByUserStr = url.searchParams.get('liked_by_user');
    const likedByUser = likedByUserStr === 'true';
    const userIdTargetStr = url.searchParams.get('user_id_target');
    const likedTodayStr = url.searchParams.get('liked_today');
    
    const apiOptions = {
        limit: 100,
        page: 1,
        sorted_by: sortedBy,
        liked_by_user: likedByUser,
        user_id_target: userIdTargetStr ? Number(userIdTargetStr) : null,
        liked_today: likedTodayStr ? Boolean(likedTodayStr) : false
    };
    
    const response = await coinderApi.coins(apiOptions);
    
    if (!response || response.data === undefined) {
        throw new Error('Failed to load data', {
            code: 500
        });
    }
    
    return [response.data, apiOptions];
}
