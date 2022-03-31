import * as React from 'react';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import { SWRConfig } from "swr";
import fetchJson from "../lib/utils/fetchJson";

const theme = createTheme({});

export default function App({Component, ...pageProps}) {
    return (
        <SWRConfig
            value={{
                fetcher: fetchJson,
                onError: (err) => {
                    console.error(err);
                },
            }}
        >
            <Component {...pageProps} />
        </SWRConfig>
    )
}
