#   e v e n t  
  
 e v e n t   i m p l e m e n t  
  
 # #   i n t e r f a c e  
  
 ` ` ` g o  
 t y p e   E v e n t   i n t e r f a c e   {  
 	 T y p e ( )   s t r i n g  
 	 S o u r c e ( )   i n t e r f a c e { }  
 	 V a l u e ( )   i n t e r f a c e { }  
 }  
 ` ` `  
  
 ` ` ` g o  
 t y p e   L i s t e n e r   f u n c ( E v e n t )  
 t y p e   D i s p a t c h e r   i n t e r f a c e   {  
 	 A d d E v e n t L i s t e n e r ( s t r i n g ,   L i s t e n e r )  
 	 R e m o v e E v e n t L i s t e n e r ( s t r i n g ,   L i s t e n e r )  
 	 D i s p a t c h E v e n t ( E v e n t )  
 }  
 ` ` `  
  
 # #   e x a m p l e  
  
 ` ` ` g o  
 c o n s t   (  
 	 a d d r e s s   =   " l o c a l h o s t : 9 0 9 0 "  
 )  
  
 f u n c   s e r v e r ( )   {  
 	 l ,   e r r   : =   n e t . L i s t e n ( " t c p " ,   a d d r e s s )  
 	 i f   e r r   ! =   n i l   {  
 	 	 p a n i c ( e r r )  
 	 }  
 	 d e f e r   l . C l o s e ( )  
  
 	 d i s p a t c h e r   : =   e v e n t . N e w E v e n t D i s p a t c h e r ( s t r u c t { } { } )  
 	 d i s p a t c h e r . A d d E v e n t L i s t e n e r ( " m s g " ,   f u n c ( e   e v e n t . E v e n t )   {  
 	 	 m s g   : =   e . V a l u e ( ) . ( [ ] b y t e )  
 	 	 f m t . P r i n t l n ( " >   " ,   s t r i n g ( m s g ) )  
 	 } )  
 	 d i s p a t c h e r . A d d E v e n t L i s t e n e r ( " l e a v e " ,   f u n c ( e   e v e n t . E v e n t )   {  
 	 	 a d d r   : =   e . V a l u e ( ) . ( n e t . A d d r )  
 	 	 f m t . P r i n t l n ( " >   n o d e   l e a v e : " ,   a d d r . S t r i n g ( ) )  
 	 } )  
  
 	 f m t . P r i n t l n ( " s e r v e r   i s   l i s t e n i n g   a t   " ,   l . A d d r ( ) . S t r i n g ( ) )  
 	 f o r   {  
 	 	 v a r   c o n n   n e t . C o n n  
 	 	 c o n n ,   e r r   =   l . A c c e p t ( )  
 	 	 i f   e r r   ! =   n i l   {  
 	 	 	 i f   e r r   = =   i o . E O F   {  
 	 	 	 	 c o n t i n u e  
 	 	 	 }  
 	 	 	 p a n i c ( e r r )  
 	 	 }  
  
 	 	 g o   f u n c ( c o n n   n e t . C o n n )   {  
 	 	 	 n   : =   0  
 	 	 	 r   : =   b u f i o . N e w R e a d e r ( c o n n )  
 	 	 	 f o r   {  
 	 	 	 	 b u f   : =   m a k e ( [ ] b y t e ,   1 0 2 4 )  
 	 	 	 	 n ,   e r r   =   r . R e a d ( b u f )  
 	 	 	 	 i f   e r r   ! =   n i l   {  
 	 	 	 	 	 i f   e r r   = =   i o . E O F   {  
 	 	 	 	 	 	 c o n t i n u e  
 	 	 	 	 	 }  
 	 	 	 	 	 d i s p a t c h e r . D i s p a t c h E v e n t ( e v e n t . N e w E v e n t ( " l e a v e " ,   c o n n . R e m o t e A d d r ( ) ) )  
 	 	 	 	 	 r e t u r n  
 	 	 	 	 }  
 	 	 	 	 d i s p a t c h e r . D i s p a t c h E v e n t ( e v e n t . N e w E v e n t ( " m s g " ,   b u f [ : n ] ) )  
 	 	 	 }  
 	 	 } ( c o n n )  
 	 }  
 }  
  
 f u n c   c l i e n t ( )   {  
 	 c o n n ,   e r r   : =   n e t . D i a l ( " t c p " ,   a d d r e s s )  
 	 i f   e r r   ! =   n i l   {  
 	 	 p a n i c ( e r r )  
 	 }  
 	 d e f e r   c o n n . C l o s e ( )  
  
 	 f m t . P r i n t l n ( " c l i e n t   d i a l   s e r v e r   s u c c e s s ,   p l e a s e   i n p u t   m e s s a g e   t o   s e n d " )  
  
 	 f o r   {  
 	 	 v a r   s   s t r i n g  
 	 	 _ ,   e r r   =   f m t . S c a n f ( " % s \ n " ,   & s )  
 	 	 i f   e r r   ! =   n i l   {  
 	 	 	 p a n i c ( e r r )  
 	 	 }  
 	 	 _ ,   _   =   c o n n . W r i t e ( [ ] b y t e ( s ) )  
 	 }  
 }  
  
 f u n c   m a i n ( )   {  
 	 m o d   : =   f l a g . S t r i n g ( " m o d " ,   " s e r v e r " ,   " s e r v e r / c l i e n t " )  
 	 f l a g . P a r s e ( )  
  
 	 i f   * m o d   = =   " s e r v e r "   {  
 	 	 s e r v e r ( )  
 	 }   e l s e   {  
 	 	 c l i e n t ( )  
 	 }  
 }  
 ` ` `  
  
 