import * as React from 'react';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import {SWRConfig} from "swr";
import fetchJson from "../lib/utils/fetchJson";

import Typography from "@mui/material/Typography";

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

export default function App({Component, ...pageProps}) {
    return (<ThemeProvider theme={theme}>
        <Typography>
            <SWRConfig
                value={{
                    fetcher: fetchJson, onError: (err) => {
                        console.error(err);
                    },
                }}
            >
                <Component {...pageProps} />
            </SWRConfig>
        </Typography>
    </ThemeProvider>)
}
