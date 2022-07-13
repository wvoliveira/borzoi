import * as React from 'react';

import '../styles/normalize.css'
import '../styles/global.css'

import "@fontsource/bebas-neue";
import "@fontsource/open-sans";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";

import Layout from '../components/Layout';


export default function App({ Component, ...pageProps }) {
    return (
        <Layout>
            <Component {...pageProps} />
        </Layout>
    )
}
