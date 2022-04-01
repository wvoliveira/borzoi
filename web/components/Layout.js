import Head from "next/head";
import Header from "../components/Header";

import Container from "@mui/material/Container";

export default function Layout({children}) {
    return (
        <>
            <Head>
                <title>Borzoi</title>
            </Head>
            <Header/>

            <main>
                <Container maxWidth="md">
                    <div>{children}</div>
                </Container>
            </main>
        </>
    );
}
