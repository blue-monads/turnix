"use client";
import { useState } from "react";
import { usePathname } from "next/navigation";

const navs = [
  {
    name: "Spaces",
    href: "/portal/pages/spaces",
  },
];

export default function PortalLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [showMenu, setShowMenu] = useState(false);
  const pathname = usePathname();

  return (
    <>
      <nav className="bg-white border border-gray-200 dark:border-gray-700 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-800 shadow">
        <div className="container flex flex-wrap justify-between items-center mx-auto">
          <a href="/pages/portal" className="flex items-center">
            <span className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">
              Turnix Portal
            </span>
          </a>

          <div className="flex gap-2 items-center ">
            <div className="flex-none h-full text-center flex items-center justify-center">
              <div className="flex space-x-3 items-center px-1">
                <div className="flex-none flex justify-center ">
                  <div className="w-8 h-8 flex ">
                    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcShta_GXR2xdnsxSzj_GTcJHcNykjVKrCBrZ9qouUl0usuJWG2Rpr_PbTDu3sA9auNUH64&usqp=CAU" alt="profile" className="shadow rounded-full object-cover" />
                  </div>
                </div>
                <div className="hidden md:block text-sm md:text-md text-black dark:text-white">John Doe</div>
              </div>
            </div>

            <div className="flex items-center">
              <button
                id="menu-toggle"
                type="button"
                className="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600 md:hidden"
                onClick={() => setShowMenu(!showMenu)}
              >
                <span className="sr-only">Open main menu</span>
                <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M4 6h16M4 12h16m-7 6h7"
                  />
                </svg>
              </button>
            </div>

          </div>
        </div>
      </nav>

      <div className="container mx-auto p-4">
        {children}
      </div>
    </>
  );
}