import { useEffect } from "react";
import Router from "next/router";
import useSWR from "swr";
import fetchJson from "./fetchJson";

export default function useUser({
  redirectTo = "",
  redirectIfFound = false,
} = {}) {
  const { data: user, mutate: mutateUser } = useSWR("/api/v1/users/me", fetchJson);

  useEffect(() => {
    // if no redirect needed, just return (example: already on /dashboard)
    // if user data not yet there (fetch in progress, logged in or not) then don't do anything yet
    if (!redirectTo || !user) {
      return;
    }

    if (
      // If redirectTo is set, redirect if error happens.
      (redirectTo && !redirectIfFound && user?.status == "error") ||
      // If redirectIfFound is also set, redirect if error not happens.
      (redirectIfFound && user?.status == "successful")
    ) {
      console.log("Push to: " + redirectTo);
      Router.push(redirectTo);
    }
  }, [user, redirectIfFound, redirectTo]);

  return { user, mutateUser };
}