import React from "react";
import Logo from "./Logo";

type TLoginRegisterLayout = {
  children: React.ReactNode;
  button: React.ReactNode;
};

export default function LoginRegisterLayout({
  children,
  button,
}: TLoginRegisterLayout) {
  return (
    <>
      <div className="container relative h-screen flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
        {button}
        <div className="relative hidden h-full flex-col bg-muted p-10 text-white dark:border-r lg:flex">
          <div className="absolute inset-0 bg-zinc-900" />
          <Logo />
          <div className="relative z-20 mt-auto">
            <blockquote className="space-y-2">
              <p className="text-lg">
                {`"O verdadeiro objetivo da educação é despertar a paixão pelo
                aprendizado e cultivar mentes curiosas."`}
              </p>
              <footer className="text-sm">Albert Einstein</footer>
            </blockquote>
          </div>
        </div>
        <div className="lg:p-8">{children}</div>
      </div>
    </>
  );
}
