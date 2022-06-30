import * as React from 'react';
import { SWRConfig } from "swr";
import fetchJson from "../lib/utils/fetchJson";

import '../styles/normalize.css'
import '../styles/global.css'

import "@fontsource/bebas-neue";
import "@fontsource/open-sans";
import "@fontsource/roboto";
import Layout from '../components/Layout';


export default function App({ Component, ...pageProps }) {
    return (
        <Layout>
            <SWRConfig
                value={{
                    fetcher: fetchJson,
                    refreshInterval: 3000,
                    revalidateIfStale: false,
                    // onError: (err) => {
                    //     console.error(err);
                    // },
                    onErrorRetry: (error, key, config, revalidate, { retryCount }) => {
                        // Never retry on 404.
                        if (error.status === 404) return

                        // Never retry for a specific key.
                        if (key === '/api/auth/check') return

                        // Only retry up to 10 times.
                        if (retryCount >= 10) return

                        // Retry after 5 seconds.
                        setTimeout(() => revalidate({ retryCount }), 5000)
                    }
                }}
            >
                <Component {...pageProps} />
            </SWRConfig>
        </Layout>
    )
}
