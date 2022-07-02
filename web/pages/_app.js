import * as React from 'react';

import '../styles/normalize.css'
import '../styles/global.css'

import "@fontsource/bebas-neue";
import "@fontsource/open-sans";
import "@fontsource/roboto";
import Layout from '../components/Layout';


export default function App({ Component, ...pageProps }) {
    return (
        <Layout>
            <Component {...pageProps} />
        </Layout>
    )
}
