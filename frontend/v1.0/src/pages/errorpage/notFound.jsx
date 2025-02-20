import { Link } from "react-router";
import Card from "../../components/card/card";
import styles from "./notfound.module.css";

export function NotFound() {
    return (
        <>
            <Card>
                <div className={styles.container}>
                    <h1 className={styles.title}>404</h1>
                    <h2 className={styles.subtitle}>Page Not Found</h2>
                    <p className={styles.description}>Sorry, we couldn't find the page you're looking for.</p>
                    <div className={styles.linkWrapper}>
                        <Link to="/" className={styles.link}>
                            Return to Home
                        </Link>
                    </div>
                    <div className={styles.imageContainer}>
                        {/* <img 
                            src="/api/placeholder/400/300"
                            alt="404 illustration"
                            className={styles.image}
                        /> */}
                    </div>
                </div>
            </Card>
        </>
    );
}