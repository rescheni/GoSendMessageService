/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        //destination: `${process.env.NEXT_PUBLIC_API_BASE_URL}/:path*`, // 使用环境变量
        destination: "http://localhost:8080/:path*", // 转发到目标地址
      },
    ];
  },
};

module.exports = nextConfig;