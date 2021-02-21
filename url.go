package ImageFetch

const Tags = "tags="
const BaseURL = "https://gelbooru.com/index.php?page=dapi&s=post&q=index&"
const RatingSafe = "rating:safe+"
const SafeTags = "solo+-underwear+-sideboob+-pov_feet+-underboob+-upskirt+-" +
	"sexually_suggestive+-ass+-bikini+-6%2Bgirls+-comic+-greyscale+-" +
	"huge_filesize+-lovestruck+-absurdres+-artificial_vagina+-cookie_%28touhou%29+-" +
	"covering_breasts+-huge_breasts+-blood+-penetration_gesture+-" +
	"animated+-audio+-webm+" + RatingSafe

const SafeTouhouQuery = BaseURL + Tags + SafeTags
