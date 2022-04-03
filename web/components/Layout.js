import Head from "next/head";
import Header from "../components/Header";

export default function Layout({children}) {
    return (
        <>
            <Head>
                <title>Borzoi</title>
            </Head>
            <Header/>

            <main>
                <br/>
                <div>{children}</div>
            </main>
        </>
    );
}
