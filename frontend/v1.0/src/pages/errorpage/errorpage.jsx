import Card from "../../components/card/card";
import styles from "./errorpage.module.css";
import { useRouteError } from "react-router";
import { Link } from "react-router";

export default function ErrorPage() {
   const error = useRouteError();

   return (
       <Card>
           <div className={styles.container}>
               <h1 className={styles.title}>{error.code || 500}</h1>
               <h2 className={styles.subtitle}>Internal Error</h2>
               <p className={styles.description}>
                   {error.message || "Something went wrong. Please try again later."}
               </p>
               <div className={styles.linkWrapper}>
                   <Link to="/" className={styles.link}>
                       Return to Home
                   </Link>
               </div>
               <div className={styles.imageContainer}>
                   {/* <img 
                       src="/api/placeholder/400/300"
                       alt="Error illustration"
                       className={styles.image}
                   /> */}
               </div>
           </div>
       </Card>
   );
}