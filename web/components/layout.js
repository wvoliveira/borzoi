import Head from 'next/head'
import Link from 'next/link'
import { default as config } from "../lib/config";

import "@fontsource/lato";
import "@fontsource/roboto";

export default function Layout({children, home}) {
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
