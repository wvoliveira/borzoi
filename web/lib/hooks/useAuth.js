import { useEffect } from "react";
import Router from "next/router";
import useSWRImmutable from 'swr/immutable'
import fetchJson from "./fetchJson";


export default function useAuth({
    redirectTo = "",
    redirectIfFound = false,
} = {}) {
    const { data: auth, mutate: mutateAuth } = useSWRImmutable("/api/auth/check", fetchJson);

    useEffect(() => {
        // if no redirect needed, just return (example: already on /dashboard)
        // if user data not yet there (fetch in progress, logged in or not) then don't do anything yet
        if (!redirectTo || !auth) {
            return;
        }

        if (
            // If redirectTo is set, redirect if error happens.
            (redirectTo && !redirectIfFound && auth?.status == "error") ||
            // If redirectIfFound is also set, redirect if error not happens.
            (redirectIfFound && auth?.status == "successful")
        ) {
            console.log("Push to: " + redirectTo);
            Router.push(redirectTo);
        }
    }, [auth, redirectIfFound, redirectTo]);

    return { auth, mutateAuth };
}