import * as React from "react";
import Layout from "../components/Layout";
import FormRegister from "../components/FormRegister";

export default function Register() {

    return (
        <Layout>
            <div className="login">
                <FormRegister/>
            </div>
            <style jsx>{`
        .login {
          max-width: 21rem;
          margin: 0 auto;
          padding: 1rem;
          border: 1px solid #ccc;
          border-radius: 4px;
        }
      `}</style>
        </Layout>
    );
}
