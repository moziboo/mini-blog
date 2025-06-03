import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { HomePage } from './pages/home/HomePage';
import { BlogPage } from './pages/blog/BlogPage';
import { PostPage } from './pages/post/PostPage';
import { AboutPage } from './pages/about/AboutPage';
import { NotFoundPage } from './pages/common/NotFoundPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/blog" element={<BlogPage />} />
        <Route path="/post/:id" element={<PostPage />} />
        <Route path="/about" element={<AboutPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
