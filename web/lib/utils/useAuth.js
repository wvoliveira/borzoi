import { useEffect } from "react";
import Router from "next/router";

import useSWR from "swr";


export default function useAuth({
    redirectTo = "",
    redirectIfFound = false,
} = {}) {
    const { data: auth, error: error, mutate: mutateAuth } = useSWR('/api/auth/check', { refreshInterval: 0 });

    useEffect(() => {
        // if no redirect needed, just return (example: already on /dashboard)
        // if user data not yet there (fetch in progress, logged in or not) then don't do anything yet
        if (!redirectTo || !auth) return;

        if (
            // If redirectTo is set, redirect if error happens.
            (redirectTo && !redirectIfFound && (error)) ||
            // If redirectIfFound is also set, redirect if error not happens.
            (redirectIfFound && !(error))
        ) {
            Router.push(redirectTo).then(r => {
            });
        }
    }, [auth, redirectIfFound, redirectTo]);

    return { auth, mutateAuth };
}