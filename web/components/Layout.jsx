import Head from "next/head";
import Header from "./Header";

export default function Layout({children}) {
    return (
        <div className="all">
            <div className="app">
                <Head>
                    <title>Borzoi</title>
                </Head>

                <Header/>

                <main>{children}</main>
            </div>
        </div>
    );
}
