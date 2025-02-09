import { useCallback } from 'react';
import styles from './price.module.css'
import { Segmented, Statistic, Avatar } from 'antd';
import { useState } from 'react';
import { ArrowDownOutlined, ArrowUpOutlined } from '@ant-design/icons';

const USD = () => (
  <span className={styles.usd}>$</span>
);

const Positive = () => (
    <Statistic
        value={11.28}
        precision={2}
        valueStyle={{
            color: 'var(--success-color)',
            marginBottom: '0.5rem',
        }}
        prefix={<ArrowUpOutlined />}
        suffix="%"
    />
)

const Negative = () => (
    <Statistic
        value={9.3}
        precision={2}
        valueStyle={{
            color: 'var(--error-color)',
            marginBottom: '0.5rem',
        }}
        prefix={<ArrowDownOutlined />}
        suffix="%"
    />
)

export default function Price({ coin }) {
    const [type, setType] = useState('Daily');

    const onChange = useCallback((value) => {
        setType(value);
    }, []);

    return (
        <>
            <div className={styles.container}>
                <div className={styles.logo}>
                    <Avatar
                        size={64}
                        src= "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png"
                    />
                </div>
                <div className={styles.statscontainer}>
                    <Statistic 
                        value={112893} 
                        suffix={<USD />}
                    />
                    <h3 className={styles.period}>{type}</h3>

                    {type === 'Daily' ? <Positive /> : <Negative />}
                </div>
            </div>
            <Segmented
                options={['Daily', 'Weekly', 'Monthly', 'Yearly']}
                onChange={onChange}
            />
        </>
    );
}