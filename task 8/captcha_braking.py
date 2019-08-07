try:
    from PIL import Image
except ImportError:
    import Image
import pytesseract

pytesseract.pytesseract.tesseract_cmd = 'C:\Program Files\Tesseract-OCR\tesseract'
k=list(pytesseract.image_to_string(Image.open('image.png')))
if "+" in k:
    print(int(k[0])+int(k[2]))
