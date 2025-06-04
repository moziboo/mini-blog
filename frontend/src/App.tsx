import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { FeedPage } from './pages/feed/FeedPage';
import { PostPage } from './pages/post/PostPage';
import { NotFoundPage } from './pages/common/NotFoundPage';
import { MainLayout } from './components/layout/MainLayout';

function App() {
  return (
    <BrowserRouter>
      <MainLayout>
        <Routes>
          <Route path="/" element={<FeedPage />} />
          <Route path="/post/:id" element={<PostPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </MainLayout>
    </BrowserRouter>
  );
}

export default App;
