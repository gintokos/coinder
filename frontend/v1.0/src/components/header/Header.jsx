import styles from './header.module.css'
import { Link } from 'react-router'

export default function Header() {
    return (
        <>
            <Link to="/">Home</Link>
            <Link to="browsing">Browsing</Link>
        </>
    )
}