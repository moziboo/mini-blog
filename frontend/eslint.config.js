import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import reactDom from 'eslint-plugin-react-dom';
import tseslint from 'typescript-eslint';
import reactX from 'eslint-plugin-react-x';
import prettierPlugin from 'eslint-plugin-prettier';

export default tseslint.config(
  { ignores: ['dist'] },
  {
    extends: [
      ...tseslint.configs.recommendedTypeChecked,
      ...tseslint.configs.stylisticTypeChecked,
      'plugin:prettier/recommended',
    ],
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
    },
    plugins: {
      'react-x': reactX,
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      'react-dom': reactDom,
      prettier: prettierPlugin,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      ...reactDom.configs.recommended.rules,
      ...reactX.configs['recommended-typescript'].rules,
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true },
      ],
      'prettier/prettier': 'warn',
    },
  }
);
