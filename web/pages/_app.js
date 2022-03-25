import * as React from 'react';
import {createTheme, ThemeProvider} from '@mui/material/styles';

import "@fontsource/lato/700.css";
import "@fontsource/roboto/700.css";

const theme = createTheme({
    // palette: {
    //     type: 'light',
    //     primary: {
    //         main: '#000814',
    //     },
    //     secondary: {
    //         main: '#ffc300',
    //     },
    //     error: {
    //         main: '#e63946',
    //     },
    // },
    // typography: {
    //     fontFamily: 'Lato, Roboto',
    // },
    // spacing: 8,
    // shape: {
    //     borderRadius: 5,
    // },
    // overrides: {
    //     MuiAppBar: {
    //         colorInherit: {
    //             backgroundColor: '#fafafa',
    //             color: '#000000',
    //         },
    //     },
    // },
    // props: {
    //     MuiAppBar: {
    //         color: 'inherit',
    //     },
    //     MuiTooltip: {
    //         arrow: true,
    //     },
    // },
});

export default function App({Component, pageProps}) {
    return <ThemeProvider theme={theme}><Component {...pageProps} /></ThemeProvider>;
}
