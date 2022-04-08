import * as React from 'react';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { SWRConfig } from "swr";
import fetchJson from "../lib/utils/fetchJson";

import "@fontsource/open-sans";
import "@fontsource/roboto";
import "@fontsource/inter";
import "@fontsource/abel"

const theme = createTheme({
    typography: {
        fontFamily: ['Open Sans', 'Roboto', 'Inter', 'Abel'].join(','),
    },
    palette: {
        type: 'light',
        primary: {
            main: '#424242',
        },
        secondary: {
            main: '#e5e5e5',
        },
    },
    spacing: 10,
    shape: {
        borderRadius: 4,
    },
    overrides: {
        MuiAppBar: {
            colorInherit: {
                backgroundColor: '#FFFFFF',
                color: '#000000',
            },
        },
    },
    props: {
        MuiAppBar: {
            color: 'inherit',
        },
    },
});

export default function App({ Component, ...pageProps }) {
    return (<ThemeProvider theme={theme}>
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
    </ThemeProvider>)
}
