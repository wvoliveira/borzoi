import React from "react";
import Layout from "../components/Layout";
import useUser from "../lib/utils/useUser";
import useEvents from "../lib/utils/useEvents";
import {Typography} from "@mui/material";
import ProfileInfo from "../components/ProfileInfo";

// Make sure to check https://nextjs.org/docs/basic-features/layouts for more info on how to use layouts
export default function Profile() {
    return (<>
            <h2>Profile</h2>
            <ProfileInfo />
        </>
    );
}
