import Link from "next/link";
import * as React from "react";
import {AuthAPI} from "../lib/api/auth";
import Router, {useRouter} from "next/router";
import { useHistory } from "react-router-dom";

export default function Logout() {
    const [isLoading, setLoading] = React.useState(false);
    const [errors, setErrors] = React.useState([]);

    const router = useRouter()
    const history = useHistory();

    const handleLogout = async (e) => {
        e.preventDefault();

        try {
            const {data, status} = await AuthAPI.Logout();
            if (status !== 200 && status !== 500) {
                setErrors(data.message);
                console.log(data.message);
            }

            if (status === 200) {
                console.log(data);
                await router.replace('/');
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    }

    return (
        <Link href={"/logout"} >
            <a onClick={(e) => handleLogout(e, "/logout")}>Logout</a>
        </Link>
    )
}
