/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://0.0.0.0:8080/:path*',
      },
    ];
  },
};

module.exports = nextConfig; 