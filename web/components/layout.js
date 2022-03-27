import * as React from "react";
import Head from 'next/head'
import Link from 'next/link'
import {default as config} from "../lib/utils/config";
import useSWR from 'swr'
import * as cookie from "../lib/utils/cookie";

import "@fontsource/lato";
import "@fontsource/roboto";
import {Fetcher} from "../lib/utils/fetcher";
import Logout from "./logout";
import {IsAuthenticated} from "../lib/hooks/auth";
import {useEffect} from "react";

export default function Layout({children, home}) {
    const [session, setSession] = React.useState();

    useEffect(() => {
        setSession(cookie.GetCookie("session"));
    }, []);

    return (
        <div>
            <Head>
                <title>{config.SiteTitle}</title>
                <link rel="icon" href="/favicon.ico"/>
                <meta
                    name="description"
                    content="Shortener app"
                />
                <meta name="og:title" content={config.SiteTitle}/>
            </Head>

            {session
                ? (
                    <div>
                        <Link href="/profile">
                            Profile
                        </Link>
                        {" | "}
                        <Logout />
                    </div>
                ) : (
                    <div>
                        <Link href="/login">
                            Login
                        </Link>
                        {" | "}
                        <Link href="/register">
                            Register
                        </Link>
                    </div>
                )}

            <main>{children}</main>
            <br/>

            {!home && (
                <div>
                    <Link href="/">
                        <a>{"< Home "}</a>
                    </Link>
                </div>
            )}
        </div>
    )
}
