import { useCallback } from 'react';
import styles from './price.module.css'
import { Segmented, Statistic, Avatar } from 'antd';
import { useState } from 'react';
import { ArrowDownOutlined, ArrowUpOutlined } from '@ant-design/icons';

const USD = () => (
  <span className={styles.usd}>$</span>
);

const PercentChange = ({value}) => {
    let color, prefix
    if (value > 0) {
        color = 'var(--success)'
        prefix = <ArrowUpOutlined />
    } else {
        color = 'var(--error)'
        prefix = <ArrowDownOutlined />
    }
    value =  Math.abs(value)
    
    return (
        <Statistic
            value={value}
            precision={2}
            valueStyle={{
                color: color,
            }}
            prefix={prefix}
            suffix="%"
        />
    )
}

const formatNumber = (num) => {
    const number = Number(num);
    
    if (number >= 0.1 && number <= 10) {
      return number.toLocaleString(undefined, {
        minimumFractionDigits: 2,
        maximumFractionDigits: 4
      });
    }
    
    if (number < 0.1) {
      return number.toExponential(4);
    }
    
    return number.toLocaleString(undefined, {
      maximumFractionDigits: 2
    });
};
    

export default function Price({ coin }) {
    const [type, setType] = useState('Daily');
    const [value, setValue] = useState(coin.percent_change_24h);

    const onChange = useCallback((value) => {
        setType(value)

        setValue(() => {
            switch (value) {
                case 'Hourly':
                    return coin.percent_change_1h
                case 'Daily':
                    return coin.percent_change_24h
                case 'Weekly':
                    return coin.percent_change_7d
            }
        })
    }, []);


    return (
        <>
            <div style={{display: 'flex', alignItems: 'center', flexDirection: 'column'}}>
                
                <div className={styles.container}>

                    <div className={styles.logo}>
                        <Avatar
                            size={64}
                            src={coin.logo}
                        />
                    </div>

                    <div className={styles.statscontainer}>
                        <Statistic 
                            value={formatNumber(coin.price)}  
                            valueStyle={{
                                color: 'var(--text-primary)',
                                opacity: 0.8,
                            }}
                            suffix={<USD />}
                            />
                        <h3 className={styles.period}>{type}</h3>

                        <PercentChange value={value} />
                    </div>
                </div>
                <div className={styles.segmentedWrapper}>
                    <Segmented
                        options={['Hourly', 'Daily', 'Weekly']}
                        onChange={onChange}
                        value={type}
                    />
                </div>

            </div>
        </>
    );
}