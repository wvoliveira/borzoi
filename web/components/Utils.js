import Link from "next/link";
import * as React from "react";
import {Typography} from "@mui/material";

export const LinkStyle = ({href, text, ...props}) => {
    return (
        <Link href={href} passHref>
            <a style={{ textDecoration: 'none', color: '#000' }} {...props}>{text}</a>
        </Link>
    )
}
