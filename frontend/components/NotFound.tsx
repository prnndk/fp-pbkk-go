import { Metadata } from "next";
export const metadata: Metadata = {
  title: "Access Forbidden",
};

export default function NotFound() {
  return (
    <div className="flex h-[85vh] items-center justify-center">
      <div className="space-y4 flex flex-col items-center gap-1">
        <h1>
          <b className="text-9xl text-white">404</b>
        </h1>
        <p className="text-2xl text-gray-500 dark:text-gray-400">
          Data cannot be found
        </p>
      </div>
    </div>
  );
}
