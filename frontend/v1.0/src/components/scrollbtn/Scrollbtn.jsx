import styles from './ScrollBtn.module.css';

export default function ScrollBtn({ onClick, className }) {
    return (
        <button 
            className={`${styles.scrollButton} ${className}`}
            onClick={onClick}
            aria-label="Scroll"
        >
            <svg 
                width="14" 
                height="14" 
                viewBox="0 0 14 14" 
                fill="none" 
                xmlns="http://www.w3.org/2000/svg"
            >
                <path 
                    d="M2 8L7 3L12 8" 
                    stroke="#1a1a1a" 
                    strokeWidth="2.5" 
                    strokeLinecap="round" 
                    strokeLinejoin="round"
                />
            </svg>
        </button>
    );
}