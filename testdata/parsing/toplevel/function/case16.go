package main
func f(a, b int, z float, opt ...interface{}) (success bool) {}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(f)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('f')
    PsiElement(()('(')
    FunctionParameterListImpl
      FunctionParameterImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('a')
        PsiElement(,)(',')
        PsiWhiteSpace(' ')
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('b')
        PsiWhiteSpace(' ')
        TypeNameImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('int')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      FunctionParameterImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('z')
        PsiWhiteSpace(' ')
        TypeNameImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('float')
      PsiElement(,)(',')
      PsiWhiteSpace(' ')
      FunctionParameterVariadicImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('opt')
        PsiWhiteSpace(' ')
        PsiElement(...)('...')
        TypeInterfaceImpl
          PsiElement(KEYWORD_INTERFACE)('interface')
          PsiElement({)('{')
          PsiElement(})('}')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    FunctionResult
      PsiElement(()('(')
      FunctionParameterListImpl
        FunctionParameterImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('success')
          PsiWhiteSpace(' ')
          TypeNameImpl
            LiteralIdentifierImpl
              PsiElement(IDENTIFIER)('bool')
      PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiElement(})('}')
