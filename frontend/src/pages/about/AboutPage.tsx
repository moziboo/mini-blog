import { MainLayout } from '../../components/layout/MainLayout';

export function AboutPage() {
  return (
    <MainLayout>
      <div className="max-w-3xl mx-auto">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">
          About MiniBlog
        </h1>

        <div className="prose prose-lg max-w-none">
          <p>
            Welcome to MiniBlog, a simple and elegant platform for sharing
            thoughts, ideas, and stories. This blog was created as a
            demonstration project to showcase modern web development practices
            and techniques.
          </p>

          <h2>Our Mission</h2>
          <p>
            At MiniBlog, we believe in the power of simplicity and the
            importance of good design. Our mission is to provide a clean and
            intuitive interface for both readers and writers, focusing on what
            matters most: the content.
          </p>

          <h2>Technology Stack</h2>
          <p>MiniBlog is built with modern technologies including:</p>
          <ul>
            <li>React for the frontend</li>
            <li>Tailwind CSS for styling</li>
            <li>Vite for fast development and optimized builds</li>
            <li>TypeScript for type safety</li>
          </ul>

          <h2>Future Plans</h2>
          <p>
            We're constantly working to improve MiniBlog. Some of our planned
            features include:
          </p>
          <ul>
            <li>User authentication and profiles</li>
            <li>Comments and interactions</li>
            <li>Rich text editing</li>
            <li>Categories and tags</li>
            <li>Dark mode support</li>
          </ul>

          <h2>Contact Us</h2>
          <p>
            Have questions, suggestions, or feedback? We'd love to hear from
            you! Reach out to us at
            <a
              href="mailto:contact@miniblog.example"
              className="text-blue-600 hover:text-blue-800 ml-1"
            >
              contact@miniblog.example
            </a>
            .
          </p>
        </div>
      </div>
    </MainLayout>
  );
}
